# FlagFairway UI

This is the Preact-based frontend for FlagFairway, built with TypeScript, Sass, and bundled using esbuild.

## Features

- ⚡️ Fast, modern UI with [Preact](https://preactjs.com/)
- ✨ TypeScript for type safety
- 🎨 Styles with Sass (`.sass`)
- ⚙️ Bundled with [esbuild](https://esbuild.github.io/)
- 📦 Static assets copied automatically to the build output

## Project Structure

```
internal/ui/
  build/           # Compiled output (bundle.js, bundle.css, index.html, etc.)
  src/
    index.tsx      # Entry point, renders the app
    page.tsx       # Main page component
    style.sass     # Sass styles
  static/          # Static assets (favicon, index.html, etc.)
  build.js         # Build script (esbuild + sass + static copy)
  package.json
  tsconfig.json
```

## Getting Started

### 1. Install dependencies

```sh
npm install
```

### 2. Build the project

```sh
npm run build
```

- This will:
  - Bundle your TypeScript/Preact code
  - Compile Sass to CSS
  - Copy static assets from `static/` to `build/`

### 3. Preview

Open `build/index.html` in your browser.

## Development Notes

- **Entry point:** [`src/index.tsx`](src/index.tsx)
- **Main component:** [`src/page.tsx`](src/page.tsx)
- **Styles:** [`src/style.sass`](src/style.sass)
- **Static assets:** Place any files you want copied to the output in [`static/`](static/)
- **Build output:** All built files are in [`build/`](build/)

## Customization

- To add more pages or components, create new `.tsx` files in `src/` and import them as needed.
- To change the theme or styles, edit `src/style.sass`.

## Troubleshooting

- If you see errors about missing font loaders (e.g. `.woff`), add the appropriate loader to your `esbuild` config in `build.js`.
- If you add new static files, re-run the build to copy them to `build/`.

## License

MIT

---