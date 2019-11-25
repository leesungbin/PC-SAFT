import React, { useState } from 'react';
import { Group, Line } from 'react-konva';
import Point from './Point';

type XY = {
  x: number,
  y: number
}
type TieProps = {
  liq: XY,
  vap: XY,
  info: { L: number[], V: number[] }
};

export default function Tie({ liq, vap, info }: TieProps) {
  const [hover, setHover] = useState(false);
  return (
    <Group>
      <Point setHover={setHover} fill="blue" info={info.L} x={liq.x} y={liq.y} hover={hover} />
      <Point setHover={setHover} fill="hotpink" info={info.V} x={vap.x} y={vap.y} hover={hover} />
      {hover && <Line points={[liq.x, liq.y, vap.x, vap.y]} stroke="green" width={1} />}
    </Group>
  );
}