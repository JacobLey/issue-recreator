# Issue - SWC file name extension

1. Open this package in devcontainer
2. Run `pnpm i`
3. Run `./node_modules/.bin/swc ./src -d ./dist --copy-files --config-file .swcrc --strip-leading-paths`
4. Inspect `/dist` dir for:
   * Wrong extension
   * Only two files omitted (the two index files write over eachother)
   * Invalid reference to `esm.mjs` file