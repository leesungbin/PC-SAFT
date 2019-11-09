export function coordChange(a: number, b: number, c: number): ({x: number, y: number}) {
  const sq3 = Math.sqrt(3);
  const y = a / 2 * sq3;
  const x = c + y / sq3;
  return { x, y };
}