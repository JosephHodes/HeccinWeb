module.exports = {
  purge: {
    enabled: true,
    content: ["./src/**/*.{html,ts}"],
  },
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      backgroundColor: {
        blu: "#0099FF",
        secondblu: "#005994",
      },
      spacing: {
        secondWave: "-20%",
      },
      height: {
        container: "400vh",
        waves: "200%",
      },
      margin: {
        proper: "3.75rem",
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
