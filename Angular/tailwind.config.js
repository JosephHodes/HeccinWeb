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
        lastblu: "#003C64",
        bottomBlack: "#00131F"
      },
      spacing: {
        secondWave: "-20%",
      },
      height: {
        container: "400vh",
        waves: "200%",
        littlecontainer: "60vh"
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
