import React from 'react';

// t : type, 0: liq, 1 : vap
export default function InfoBox({ liq, vap, x, y }: { liq: number[], vap: number, x: number, y: number}) {
  return (
    <mesh visible position={[x, y, 0.1]}>
      <boxGeometry attach="geometry" args={[0.007, 4, 4]} />
      <meshBasicMaterial attach="material" color="white" />
    </mesh>
  )
}