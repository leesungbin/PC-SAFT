export function coordChange(a: number, b: number, c: number): ({ x: number, y: number }) {
  const sq3 = Math.sqrt(3);
  const y = a / 2 * sq3;
  const x = c + y / sq3;
  return { x, y };
}

// points: 삼각형의 각 좌표 배열 [x1, y1, x2, y2, x3, y3]
//   2
// 1   3
export function xyTransform(x: number, y: number, points: number[], ternP: number[]): ({x: number, y: number}) {
  const comy = points[3] - points[5]; // y2-y3
  const comx = points[2] - points[4]; // x2-x3
  const p = ternP[0] * comy;
  const q = ternP[1] * comy;
  const Y = points[1] + p;
  // y=comy/comx*(x-x2)+y2
  // y=comy/comx*x-comy/comx*x2+y2
  // comx*y-comy*x+comy*x2-y2*comx = 0
  // Math.abs( comx*Y-comy*X+comy*x2-y2*comx )/Math.sqrt(comx*comx+comy*comy) == q)
  // (1) comx*Y-comy*X+comy*x2-y2*comx = q*Math.sqrt(comx*comx+comy*comy)
  // (2) comx*Y-comy*X+comy*x2-y2*comx  = -q*Math.sqrt(comx*comx+comy*comy)
  const X1 = -(q * Math.sqrt(comx * comx + comy * comy) - comx * Y - comy * points[2] + points[3] * comx) / (comy);
  const X2 = (q * Math.sqrt(comx * comx + comy * comy) - comx * Y - comy * points[2] + points[3] * comx) / (comy);
  const line = (x: number) => { return comy / comx * (x - points[2]) + points[3] };
  if (line(X1) < Y) return {x: X1, y: Y};
  return {x: X2, y: Y};
}