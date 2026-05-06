package text

import (
	"html"
	"regexp"
	"strings"
	"unicode/utf8"
)

// SanitizeText sanitizes text similar to WordPress sanitize_textarea_field()
// keepNewlines: true preserves line breaks, false collapses all whitespace
func SanitizeText(str interface{}, keepNewlines bool) string {
	// Return empty string for arrays or objects (maps/slices in Go)
	switch str.(type) {
	case map[string]interface{}, []interface{}:
		return ""
	}

	// Convert to string
	var s string
	switch v := str.(type) {
	case string:
		s = v
	default:
		if str == nil {
			return ""
		}
		s = ""
	}

	// Check for invalid UTF-8
	filtered := checkInvalidUTF8(s, false)
	if filtered == "" && s != "" {
		return ""
	}

	// Handle HTML tags if present
	if strings.Contains(filtered, "<") {
		filtered = preKsesLessThan(filtered)
		filtered = stripAllTags(filtered, false)
		filtered = strings.ReplaceAll(filtered, "<\n", "&lt;\n")
	}

	// Remove extra whitespace if newlines not needed
	if !keepNewlines {
		whitespaceRegex := regexp.MustCompile(`[\r\n\t ]+`)
		filtered = whitespaceRegex.ReplaceAllString(filtered, " ")
	} else {
		// Preserve newlines but replace tabs with spaces
		filtered = strings.ReplaceAll(filtered, "\t", " ")
	}

	filtered = strings.TrimSpace(filtered)

	// Remove percent-encoded characters (except %20 which becomes space)
	percentRegex := regexp.MustCompile(`%[a-fA-F0-9]{2}`)

	if percentRegex.MatchString(filtered) {
		// First, replace %20 with space
		filtered = strings.ReplaceAll(filtered, "%20", " ")

		// Remove all other percent encodings
		filtered = percentRegex.ReplaceAllString(filtered, "")

		// Clean up extra spaces
		spaceRegex := regexp.MustCompile(` +`)
		filtered = spaceRegex.ReplaceAllString(filtered, " ")
		filtered = strings.TrimSpace(filtered)
	}

	// Escape HTML special characters in remaining text (for unclosed tags)
	filtered = strings.ReplaceAll(filtered, "<", "&lt;")
	filtered = strings.ReplaceAll(filtered, ">", "&gt;")

	return filtered
}

// checkInvalidUTF8 checks if string is valid UTF-8
func checkInvalidUTF8(text string, strip bool) string {
	if len(text) == 0 {
		return ""
	}

	if isValidUTF8(text) {
		return text
	}

	if strip {
		return scrubUTF8(text)
	}
	return ""
}

// preKsesLessThan handles less-than signs before KSES
func preKsesLessThan(content string) string {
	var result strings.Builder
	i := 0
	n := len(content)

	for i < n {
		if content[i] == '<' {
			// Look for closing '>'
			j := i + 1
			foundGt := false
			for j < n {
				if content[j] == '>' {
					foundGt = true
					j++
					break
				}
				j++
			}

			if foundGt {
				// Valid HTML tag with closing >, keep it as is for now
				result.WriteString(content[i:j])
			} else {
				// No closing > found, keep the < as is (will be escaped later)
				result.WriteByte('<')
				i++
				continue
			}
			i = j
		} else {
			result.WriteByte(content[i])
			i++
		}
	}

	return result.String()
}

// stripAllTags removes all HTML tags
func stripAllTags(text string, removeBreaks bool) string {
	if text == "" {
		return ""
	}

	// Remove script tags and their content completely
	scriptRegex := regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`)
	text = scriptRegex.ReplaceAllString(text, "")

	// Remove style tags and their content completely
	styleRegex := regexp.MustCompile(`(?i)<style[^>]*>.*?</style>`)
	text = styleRegex.ReplaceAllString(text, "")

	// Remove all HTML tags, but preserve text between them
	// This also removes comments
	tagRegex := regexp.MustCompile(`<[^>]*>`)
	text = tagRegex.ReplaceAllString(text, "")

	// Normalize whitespace: ensure single spaces between words
	spaceRegex := regexp.MustCompile(`\s+`)
	text = spaceRegex.ReplaceAllString(text, " ")

	// Preserve intentional spaces - don't collapse all spaces
	// Just normalize multiple spaces to single space

	return strings.TrimSpace(text)
}

// isUTF8Charset checks if charset is UTF-8
func isUTF8Charset(charsetSlug string) bool {
	charsetSlug = strings.ToUpper(charsetSlug)
	return charsetSlug == "UTF-8" || charsetSlug == "UTF8"
}

// isValidUTF8 validates UTF-8 string
func isValidUTF8(bytes string) bool {
	return utf8.ValidString(bytes)
}

// scrubUTF8 removes invalid UTF-8 characters
func scrubUTF8(text string) string {
	if utf8.ValidString(text) {
		return text
	}

	var result strings.Builder
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		if r == utf8.RuneError && size == 1 {
			// Invalid byte, replace with Unicode replacement character
			result.WriteRune('\uFFFD')
			i++
		} else {
			result.WriteRune(r)
			i += size
		}
	}
	return result.String()
}

// escapeHTML escapes HTML special characters
func escapeHTML(text string) string {
	safeText := checkInvalidUTF8(text, false)
	if safeText == "" && text != "" {
		return ""
	}
	return specialChars(safeText, "ENT_QUOTES", "UTF-8", false)
}

// specialChars converts special characters to HTML entities
func specialChars(text string, quoteStyle string, charset string, doubleEncode bool) string {
	if len(text) == 0 {
		return ""
	}

	// Check if we need to do any conversion
	needsConversion := false
	for _, ch := range text {
		if ch == '&' || ch == '<' || ch == '>' || ch == '"' || ch == '\'' {
			needsConversion = true
			break
		}
	}

	if !needsConversion {
		return text
	}

	charset = canonicalCharset(charset)
	_ = charset // For future use

	// Normalize entities if not double encoding
	if !doubleEncode {
		text = normalizeEntities(text, "html")
	}

	// First escape the string using Go's html.EscapeString
	// This escapes: &, <, >, " (and ' for XML but not for HTML)
	escaped := html.EscapeString(text)

	// Handle single quotes if needed (ENT_QUOTES in PHP)
	if quoteStyle == "ENT_QUOTES" || quoteStyle == "single" {
		// Replace ' with &#039; (same as WordPress)
		escaped = strings.ReplaceAll(escaped, "'", "&#039;")
	} else if quoteStyle == "double" {
		// For double quotes, we use the default behavior (already escaped by html.EscapeString)
		// Do nothing extra
	}
	// For ENT_NOQUOTES, quotes are not escaped (html.EscapeString already escapes double quotes)
	// We need to unescape double quotes if needed
	if quoteStyle == "ENT_NOQUOTES" {
		escaped = strings.ReplaceAll(escaped, "&quot;", "\"")
	}

	return escaped
}

// canonicalCharset returns canonical charset name
func canonicalCharset(charset string) string {
	if isUTF8Charset(charset) {
		return "UTF-8"
	}

	upper := strings.ToUpper(charset)
	if upper == "ISO-8859-1" || upper == "ISO8859-1" {
		return "ISO-8859-1"
	}

	return charset
}

// normalizeEntities normalizes HTML entities
func normalizeEntities(content string, context string) string {
	// Replace & with &amp;
	content = strings.ReplaceAll(content, "&", "&amp;")

	// Handle numeric entities (decimal)
	decimalRegex := regexp.MustCompile(`&amp;#(0*[1-9][0-9]{0,6});`)
	content = decimalRegex.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the number and convert
		return normalizeEntities2(match)
	})

	// Handle numeric entities (hex)
	hexRegex := regexp.MustCompile(`&amp;#[Xx](0*[1-9A-Fa-f][0-9A-Fa-f]{0,5});`)
	content = hexRegex.ReplaceAllStringFunc(content, func(match string) string {
		return normalizeEntities3(match)
	})

	// Handle named entities
	namedRegex := regexp.MustCompile(`&amp;([A-Za-z]{2,8}[0-9]{0,2});`)
	if context == "xml" {
		content = namedRegex.ReplaceAllStringFunc(content, func(match string) string {
			return xmlnamedEntities(match)
		})
	} else {
		content = namedRegex.ReplaceAllStringFunc(content, func(match string) string {
			return namedEntities(match)
		})
	}

	return content
}

// normalizeEntities2 handles decimal entities
func normalizeEntities2(match string) string {
	// Extract the numeric value
	re := regexp.MustCompile(`&amp;#([0-9]+);`)
	matches := re.FindStringSubmatch(match)
	if len(matches) > 1 {
		// Return as regular numeric entity
		return "&#" + matches[1] + ";"
	}
	return match
}

// normalizeEntities3 handles hex entities
func normalizeEntities3(match string) string {
	// Extract the hex value
	re := regexp.MustCompile(`&amp;#[Xx]([0-9A-Fa-f]+);`)
	matches := re.FindStringSubmatch(match)
	if len(matches) > 1 {
		// Return as hex entity
		return "&#x" + matches[1] + ";"
	}
	return match
}

// xmlnamedEntities handles XML named entities
func xmlnamedEntities(match string) string {
	// Map of XML named entities
	entities := map[string]string{
		"amp":  "&amp;",
		"apos": "&apos;",
		"gt":   "&gt;",
		"lt":   "&lt;",
		"quot": "&quot;",
	}

	// Extract entity name
	entityName := strings.TrimPrefix(match, "&amp;")
	entityName = strings.TrimSuffix(entityName, ";")

	if val, ok := entities[entityName]; ok {
		return val
	}
	return match
}

// namedEntities handles HTML named entities
func namedEntities(match string) string {
	// Common HTML entities map to their entity format
	entities := map[string]string{
		"nbsp":   "&nbsp;",
		"amp":    "&amp;",
		"lt":     "&lt;",
		"gt":     "&gt;",
		"quot":   "&quot;",
		"apos":   "&apos;",
		"copy":   "&copy;",
		"reg":    "&reg;",
		"euro":   "&euro;",
		"pound":  "&pound;",
		"yen":    "&yen;",
		"sect":   "&sect;",
		"deg":    "&deg;",
		"plusmn": "&plusmn;",
	}

	// Extract entity name
	entityName := strings.TrimPrefix(match, "&amp;")
	entityName = strings.TrimSuffix(entityName, ";")

	if val, ok := entities[entityName]; ok {
		return val
	}
	return match
}
