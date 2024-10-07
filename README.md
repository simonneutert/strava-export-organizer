# Strava Export Organizer

The purpose of this project is to organize your activities data in your Strava export.

After running this tool you should have your activity files sorted by the type of activity and the year it was performed in. 

The script will create a directory tree like this:

```text
./export_mapped/<:activity_typeA>/<:activity_typeA>_<:year1>/
./export_mapped/<:activity_typeA>/<:activity_typeA>_<:year2>/
./export_mapped/<:activity_typeB>/<:activity_typeB>_<:year3>/
./export_mapped/<:activity_typeC>/<:activity_typeC>_<:year1>/
./export_mapped/<:activity_typeC>/<:activity_typeC>_<:year2>/
./export_mapped/<:activity_typeC>/<:activity_typeC>_<:year3>/
./export_mapped/<:activity_typeC>/<:activity_typeD>_<:year2>/
# and so on...
```

#### Web/Browser/Docker version 🚀

https://github.com/simonneutert/strava-export-organizer-web

### Supported Languages

- en
- de

### Supported Operating Systems

- Linux 🐧
- MacOS 🍎

## Usage

- Request your Strava export from [Strava](https://www.strava.com/).
- Unzip the downloaded file.
- Download this program, for your platform and OS.
- Copy the binary into/inside the unzipped Strava export folder.
  - Say your unzipped export is here: `home/username/Downloads/my_strava_export/export_1234567890/`.
  - The strava_export_organizer binary should be: 
    `home/username/Downloads/my_strava_export/export_1234567890/strava_export_organizer.exe`
    (`.exe` was added for clarification).
- Navigate (`cd`) to the directory the binary is now in.
- Run the executable passing your language. Currently supported languages: `en`/`de`.
  - `$ ./strava_export_organizer en`
- A new folder will be created called `export_mapped/` within your export containing the organized activities.
- Enjoy the warm and fuzzy feeling of a well-organized Strava export 🧘.

## Add your language? Help developing?

See [README.Develop.md](./README.Develop.md) for more information on how to develop or i.e. add your language to the project.
