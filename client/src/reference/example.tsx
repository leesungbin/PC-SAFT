import React, { useRef } from 'react';
import { Canvas, useFrame } from 'react-three-fiber';

const App: React.FC = () => {
  return (
    <div className="App">
      <Canvas>
        <Thing />
      </Canvas>
    </div>
  );
}


export function Thing() {
  const ref = useRef<any>()
  useFrame(() => [(ref.current.rotation.x += 0.008), (ref.current.rotation.y += 0.008)])
  return (
    <mesh
      ref={ref}
      onClick={e => console.log('click')}
      onPointerOver={e => console.log('hover')}
      onPointerOut={e => console.log('unhover')}>
      {/* <ambientLight intensity={0.5} />
      <spotLight intensity={0.5} position={[1, 0, 0]} /> */}
      <boxBufferGeometry attach="geometry" args={[0, 0, 0]} />
      <meshNormalMaterial attach="material" />
      {/* <lineBasicMaterial/> */}
    </mesh>
  )
}
export default App;
