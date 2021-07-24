module.exports = {
  // mode: 'jit',
  purge: ['./src/**/*.{html,ts}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    colors: {
      blue: {
        brightBlue: "#0099FF"
      }

    },
    extend: {
      backgroundImage: background => ({
        'smallWaves': "url('assets/Media/smallwaves.svg')",
        'middleWaves': "url('assets/Media/middle.svg')",
        'largeWaves': "url('assets/Media/bigwaves.svg')"
      })
    },
  },
  variants: {
    extend: {}
    ,
  }
  ,
  plugins: [require('@tailwindcss/aspect-ratio')
    , require('@tailwindcss/forms')
    , require('@tailwindcss/line-clamp')
    , require('@tailwindcss/typography')
  ],
}
;
