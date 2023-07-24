/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'background': '#05050A',
        'selection': '#44475a',
        'hover': '#212024',
        'border': '#23252C',
        'text-primary': '#fdfeffec',
        'text-secondary': '#aaaaaa',
      }
    },
  },
  plugins: [require('@tailwindcss/typography')],
}

