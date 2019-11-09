import { useFrame, Canvas, ReactThreeFiber } from 'react-three-fiber';
import { Mesh } from 'three';
import { useRef } from 'react';
import React from 'react';

export default function point() {
  const ref = useRef<Mesh>(null);
  useFrame(() => ([ref.current!.rotation.x += 0.1, ref.current!.rotation.y += 0.1]));
  return (
    <mesh visible ref={ref} position={[0, 0, 0]} rotation={[0, 0, 0]}>
      <sphereGeometry />
    </mesh>
  )
}