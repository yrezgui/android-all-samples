package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Sample struct {
	AndroidDefaultAppID string   `json:"androidDefaultAppId"`
	CompileSdk          string   `json:"compileSdk"`
	DeclaredPermissions []string `json:"declaredPermissions"`
	Dependencies        []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"dependencies"`
	GradlePlugins []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"gradlePlugins"`
	MinSdk                   string   `json:"minSdk"`
	ModuleNames              []string `json:"moduleNames"`
	SingleActivity           bool     `json:"singleActivity"`
	TargetSdk                string   `json:"targetSdk"`
	UseAidl                  bool     `json:"useAidl"`
	UseBuildConfig           bool     `json:"useBuildConfig"`
	UseCoreLibraryDesugaring bool     `json:"useCoreLibraryDesugaring"`
	UseDataBinding           bool     `json:"useDataBinding"`
	UseFragment              bool     `json:"useFragment"`
	UseJetpackCompose        bool     `json:"useJetpackCompose"`
	UsePrefab                bool     `json:"usePrefab"`
	UseProtocolBuffers       bool     `json:"useProtocolBuffers"`
	UseRenderScript          bool     `json:"useRenderScript"`
	UseShaders               bool     `json:"useShaders"`
	UseSpotless              bool     `json:"useSpotless"`
	UseViewBinding           bool     `json:"useViewBinding"`
}

var GeminiSampleDetails = &genai.Schema{
	Type:     genai.TypeObject,
	Required: []string{"androidDefaultAppId", "minSdk", "compileSdk", "targetSdk", "moduleNames", "declaredPermissions", "singleActivity", "useFragment", "dependencies", "gradlePlugins", "useSpotless", "useJetpackCompose", "useViewBinding", "useDataBinding", "useProtocolBuffers", "useBuildConfig", "useAidl", "usePrefab", "useRenderScript", "useShaders", "useCoreLibraryDesugaring"},
	Properties: map[string]*genai.Schema{
		"androidDefaultAppId": {
			Type: genai.TypeString,
		},
		"minSdk": {
			Type: genai.TypeString,
		},
		"compileSdk": {
			Type: genai.TypeString,
		},
		"targetSdk": {
			Type: genai.TypeString,
		},
		"moduleNames": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeString,
			},
		},
		"declaredPermissions": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeString,
			},
		},
		"singleActivity": {
			Type: genai.TypeBoolean,
		},
		"useFragment": {
			Type: genai.TypeBoolean,
		},
		"dependencies": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type:     genai.TypeObject,
				Required: []string{"name", "version"},
				Properties: map[string]*genai.Schema{
					"name": {
						Type: genai.TypeString,
					},
					"version": {
						Type: genai.TypeString,
					},
				},
			},
		},
		"gradlePlugins": {
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type:     genai.TypeObject,
				Required: []string{"name", "version"},
				Properties: map[string]*genai.Schema{
					"name": {
						Type: genai.TypeString,
					},
					"version": {
						Type: genai.TypeString,
					},
				},
			},
		},
		"useSpotless": {
			Type: genai.TypeBoolean,
		},
		"useJetpackCompose": {
			Type: genai.TypeBoolean,
		},
		"useViewBinding": {
			Type: genai.TypeBoolean,
		},
		"useDataBinding": {
			Type: genai.TypeBoolean,
		},
		"useProtocolBuffers": {
			Type: genai.TypeBoolean,
		},
		"useBuildConfig": {
			Type: genai.TypeBoolean,
		},
		"useAidl": {
			Type: genai.TypeBoolean,
		},
		"usePrefab": {
			Type: genai.TypeBoolean,
		},
		"useRenderScript": {
			Type: genai.TypeBoolean,
		},
		"useShaders": {
			Type: genai.TypeBoolean,
		},
		"useCoreLibraryDesugaring": {
			Type: genai.TypeBoolean,
		},
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro")
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = GeminiSampleDetails

	// Read the file content and split into lines
	sampleList, err := os.ReadFile("./sample-list.txt")
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	sampleFolders := strings.Split(string(sampleList), "\n")

	// Remove any empty lines from lines
	sampleCount := len(sampleFolders)

	// Create a new file for the output
	outputFile, err := os.Create("./sample-data.jsonl")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer outputFile.Close()

	// Iterate through the lines
	for index, line := range sampleFolders {
		if line == "" {
			continue
		}
		codebase, err := packSampleCodebase(line)
		if err != nil {
			log.Fatalf("Error packing codebase: %v", err)
		}

		opts := &genai.UploadFileOptions{MIMEType: "text/plain"}

		file, err := client.UploadFile(ctx, "", strings.NewReader(codebase), opts)
		if err != nil {
			log.Fatalf("Error uploading file: %v", err)
		}
		defer client.DeleteFile(ctx, file.Name)

		resp, err := model.GenerateContent(ctx,
			genai.Text("Fill the project details based on the files attached to the context"),
			genai.FileData{URI: file.URI})
		if err != nil {
			log.Fatalf("Error calling the model: %v", err)
		}

		for _, part := range resp.Candidates[0].Content.Parts {
			jsonResult := fmt.Sprintf("%v\n", part)

			var sample Sample
			if err := json.Unmarshal([]byte(jsonResult), &sample); err != nil {
				panic(err)
			}

			// Append JSON output to sample-data.jsonl
			rawJson, _ := json.Marshal(sample)
			if _, err := outputFile.WriteString(string(rawJson) + "\n"); err != nil {
				panic(err)
			}
		}

		fmt.Printf("%s: %d/%d\n", line, index+1, sampleCount) // Progress indicator
	}
}

// packSampleCodebase appends file contents from a folder (including all its subfolders)
// to a string and returns it, with improved efficiency for large codebases.
func packSampleCodebase(sampleFolder string) (string, error) {
	var output strings.Builder

	// Define the file extensions and names to include
	extensions := []string{".kt", ".java", ".proto", ".kts"}
	filenames := []string{"AndroidManifest.xml", "build.gradle", "gradle.properties", "libs.versions.toml"}

	// Walk the directory tree
	err := filepath.WalkDir(sampleFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Check if the file matches the criteria
		if shouldIncludeFile(path, extensions, filenames) {
			err = appendFileContent(&output, path)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return "", err // Return error instead of panicking
	}

	return output.String(), nil
}

// appendFileContent appends the content of a file to the provided string builder.
func appendFileContent(output *strings.Builder, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	output.WriteString(strings.Repeat("=", 80))
	output.WriteString(fmt.Sprintf("\nFilepath: %s\n", path))
	output.WriteString(strings.Repeat("=", 80))
	output.WriteString(fmt.Sprintf("%s\n\n", strings.Repeat("=", 80)))
	// output.WriteString(fmt.Sprintf("%s\nFilepath: %s\n%s\n", strings.Repeat("=", 80), path, strings.Repeat("=", 80)))

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			return err
		}
		output.WriteString(line)
	}

	output.WriteString("\n\n")
	return nil
}

// shouldIncludeFile checks if the file should be included based on its extension or name
func shouldIncludeFile(path string, extensions []string, filenames []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	for _, filename := range filenames {
		if filepath.Base(path) == filename {
			return true
		}
	}
	return false
}
