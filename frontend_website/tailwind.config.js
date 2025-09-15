/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#effaf5',
          100: '#dff6eb',
          200: '#bdebd7',
          300: '#98e0c3',
          400: '#5dc99d',
          500: '#38b280',
          600: '#298d68',
          700: '#227154',
          800: '#1f5b46',
          900: '#1b4b3b',
          950: '#0d281f',
        },
      },
    },
  },
  plugins: [],
  // Force Tailwind to generate all utilities for debugging
  experimental: {
    optimizeUniversalDefaults: false,
  }
}