package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

type YAMLStruct struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		url, found := pathsToUrls[path]
		if found {
			//http.Redirect(w, r, url, StatusCode)
			// TODO: Lookup status codes
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	/*
		return func(w http.ResponseWriter, r *http.Request) {
			fallback.ServeHTTP(w, r)
		}, nil*/
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]YAMLStruct, error) {
	var ymlMap []YAMLStruct

	err := yaml.Unmarshal(yml, &ymlMap)
	if err != nil {
		return nil, err
	}
	return ymlMap, nil
}

func buildMap(parsedYaml []YAMLStruct) map[string]string {
	pathMap := make(map[string]string)
	for _, value := range parsedYaml {
		pathMap[value.Path] = value.URL
	}
	return pathMap
}
