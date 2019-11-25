import React from 'react';
type LineProps = {
  scale: number,
  lineWidth?: number,
  // transformOrigin
  p1: Point,
  p2: Point
}
export type Point = {
  x: number,
  y: number,
}
export function Line({ scale, lineWidth = 1, p1, p2 }: LineProps) {
  const a=scale*0.84;
  const angle=Math.atan((p2.y-p1.y)/(p2.x-p1.x));
  const left = 0.08*scale+p1.x*a;
  const top = 0*scale+p1.y*a;
  // console.log(left, bottom);
  // console.log(Math.tan(Math.abs(angle)));
  return (
    <div
      style={{
        position: 'absolute',
        width: a,
        backgroundColor: 'black',
        height: lineWidth,
        transform: `rotate(${angle}rad)`,
        transformOrigin: '0% 0%',
        left,
        top,
      }}
    />
  );
}