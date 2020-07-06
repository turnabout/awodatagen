# awodatagen

Simple commandline tool used to generate [https://github.com/turnabout/AWO](AWO)'s sprite sheet and game data JSON file.

Raw sprites and other JSON files located in the `assets` directory are processed and combined to create a single sprite sheet image and a game data JSON file, which are output at whichever locations are pointed at by the `AWO_SPRITESHEET` and `AWO_JSON` environment variables.
