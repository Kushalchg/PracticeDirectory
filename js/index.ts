const TextColor = (bgc: string) => {
  const bg = bgc.substring(1)
  console.log("bg color", bg)

  const r = parseInt(bg.substring(0, 2), 16)
  const g = parseInt(bg.substring(2, 4), 16)
  const b = parseInt(bg.substring(4, 6), 16)

  const userRelativeLuminance = 0.2126 * r + 0.7152 * g + 0.0722 * b
  const contrastRatio = userRelativeLuminance / 255
  console.log("contrast Ratio", contrastRatio)

  return contrastRatio >= 0.5 ? "#000" : "#fff";
}
console.log(TextColor("#3a8d99"))
