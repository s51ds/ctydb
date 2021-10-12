package parser

import "regexp"

///////////////
// regex
var regexPfx *regexp.Regexp = regexp.MustCompile(`[0-9A-Za-z/]{1,}`)
var regexOverrideCqZone *regexp.Regexp = regexp.MustCompile(`\([0-9]{1,2}\)`)
var regexOverrideItuZone *regexp.Regexp = regexp.MustCompile(`[([0-9]{1,2}]`)
var regexZone *regexp.Regexp = regexp.MustCompile(`[0-9]{1,2}`)
var regexOverrideLatLon *regexp.Regexp = regexp.MustCompile(`<[0-9-]+\.[0-9]+/[0-9-]+\.[0-9]+>`)
var regexOverrideTimeOffset *regexp.Regexp = regexp.MustCompile(`~[0-9-.]+~`)
var regexIsComment *regexp.Regexp = regexp.MustCompile(`^#`)
