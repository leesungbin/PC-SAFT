import React from 'react';
import { Circle, Group, Text } from 'react-konva';

type PointProps = {
  setHover: (state: boolean) => void,
  fill: string,
  x: number,
  y: number,
  hover: boolean,
  info: number[],
};
export default function Point({ setHover, fill, x, y, info, hover }: PointProps) {
  return (
    <Group>
      <Circle onMouseEnter={() => setHover(true)} onMouseOut={() => setHover(false)} radius={3} fill={fill} x={x} y={y} />
      {hover &&
        <Text text={
          JSON.stringify(info, (key, val) => { return val.toFixed ? Number(val.toFixed(3)) : val })}
          x={x} y={y} width={120}  align="center" fill="black" draggable />}
    </Group>
  )
}