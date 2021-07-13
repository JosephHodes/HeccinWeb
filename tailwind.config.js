module.exports = {
  purge: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
    backgroundColor: (theme) => ({
      ...theme("colors"),
      background: "#0A192E",
    }),
    gradientColorStops: (theme) => ({
      "h-primary": "#629E92",
      "h-secondary": "#79DDCA",
    }),
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
