import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';

class Home extends React.Component {

  render() {
    return (
      <Content>
        <h1>Home</h1>
        <Canvas>
          <mesh visible userData={{ test: 'hello' }} position={[0, 0, 0]} rotation={[0, 0, 0]}>
            <sphereGeometry attach="geometry" args={[1, 16, 16]} />
            {/* <meshStandardMaterial attach="material" color="#aaaaaa" transparent /> */}
            <meshNormalMaterial attach="material" />
          </mesh>
          <mesh visible userData={{ test: 'hello' }} position={[0, 2, 1]} rotation={[0, 0, 0]}>
            <sphereGeometry attach="geometry" args={[1, 16, 16]} />
            {/* <meshStandardMaterial attach="material" color="#abcdef" transparent /> */}
            <meshNormalMaterial attach="material" />
          </mesh>
        </Canvas>
      </Content>
    );
  }
}
export default Home;