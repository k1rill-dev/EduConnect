const flowbite = require("flowbite-react/tailwind");

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./src/**/*.{js,jsx,ts,tsx}", // Adjust this path to match your React components
        flowbite.content(),
        // "./node_modules/@material-tailwind/react/components/**/*.{js,ts,jsx,tsx}",
        // "./node_modules/@material-tailwind/react/theme/components/**/*.{js,ts,jsx,tsx}",

    ],
    theme: {
        extend: {
            colors: {
                'gosu-blue': '#007bff',
                'gosu-gray': '#f8f9fa',
                'gosu-dark-gray': '#6c757d',
            },
        },
    },
    plugins: [
        flowbite.plugin()
    ],
};