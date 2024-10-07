# Developer Readme

I am happy to accept contributions to this project. If you want to add a new language, fix a bug or add a new feature, you are welcome to do so. Best to open an issue first, so we can discuss the changes you planned upfront.

## Requirements

- Golang 1.22 or higher

## Add a new language

I will keep this short and sweet. And show you how the current implementation of languages is done.

### Example "english/de"

See files in `src/en/en.go` and `src/de/de.go` first.

Also tests in `src/en/en_test.go` and `src/de/de_test.go`.

1. Create a new directory in `src/` with the language code, i.e. `es` for Spanish.
2. Create a new file in the new directory, i.e. `es.go`.
3. Copy the content of `src/en/en.go` into the new file.
4. Add the language code to the `languageConfig` map in `src/main.go` (in the `init` function).
5. Add support for the new language in the `getLanguage` function in `src/main.go`.
  ```go
    func mapData(records [][]string, language string) []map[string]string {
        var result []map[string]string
        if language == "de" {
            result = mapGenericRows(records, language, de.MapDateDe, de.MapFilenameDe)
        }
        if language == "en" {
            result = mapGenericRows(records, language, en.MapDateEn, en.MapFilenameEn)
        }
        return result
    }
  ```
1. Copy your unzipped strava export directory into the root of the project.
2. Rename the export to `export_1`
3. Add your language. Example: `src/es/es.go` and `src/es/es_test.go`
4. Add tests for your language in `es_test.go` - I can help your with that, if you provide me a sample of the strava export in your language. The first 4 columns of your exports `activities.csv` should be sufficient.
5. You can test run the code against your test data and check the results in `export_mapped`, simply by running: `$ just test en` (replace `en` with your language, of course).
6. Have a drink 🧉 and celebrate your contribution.

#### VS Code users

Change the args in the launch.json to the parameters you need to test.

The first and single parameter is the language code, like `en` or `de` of hopefully soon `es` and `fr`.

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "args": ["en"]
        }
    ]
}
```

#### Non VS Code 

Change the default values of the `main.go` file to the parameters you need to test.

```go
language := "de"
```

`fr` for French, `es` for Spanish, and so on, then remove what's blocking you from running/testing.