const initial = require("express");
const router = initial();
const port = 8000;
router.get("/", (req, res) => {
  res.send("Hello World!");
});

router.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
