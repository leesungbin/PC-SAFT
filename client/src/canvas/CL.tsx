// ## connect line
import React from 'react';
import { Line, Group } from 'react-konva';
type CLProps = {
  x: number[],
  y: number[],
}
export function CL({ x, y }: CLProps) {
  let points = [];
  for (let i = 0; i < x.length; i++) {
    points.push(x[i]);
    points.push(y[i]);
  }
  return (
    <Line points={points} stroke="black" strokeWidth={1} />
  )
}


type CLsProps = {
  datas: { liq: { x: number, y: number }, vap: { x: number, y: number } }[]
}

export function CLs({ datas }: CLsProps) {
  let x1: number[] = [], y1: number[] = [];
  let x2: number[] = [], y2: number[] = [];

  for (let i = 0; i < datas.length; i++) {
    x1.push(datas[i].liq.x);
    y1.push(datas[i].liq.y);

    x2.push(datas[i].vap.x);
    y2.push(datas[i].vap.y);
  }

  return (
    <Group>
      <CL x={x1} y={y1} />
      <CL x={x2} y={y2} />
    </Group>
  )
}
