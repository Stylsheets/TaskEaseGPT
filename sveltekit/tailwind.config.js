/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'background': '#0F1111',
        'accent': '#161A1A',
        'primary': '#20C2DF',
        'secondary': '#144B54',
        'text': '#F2F2F2',
        'light': "rgba(255, 255, 255, 0.5)"

      }
    },
  },
  plugins: [require('@tailwindcss/typography')],
}

