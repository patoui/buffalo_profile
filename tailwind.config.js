// tailwind.config.js
module.exports = {
    purge: [
        './templates/**/*.html',
        './assets/**/*.js',
        './assets/**/*.vue',
    ],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {},
    },
    variants: {},
    plugins: [],
}
