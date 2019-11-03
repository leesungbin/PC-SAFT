import React from 'react';
import { Content } from '../components/Content';
import { Canvas } from 'react-three-fiber';
import { Thing } from '../reference/example';

class Home extends React.Component {
  render() {
    return (
      <Content>
        <h1>Home</h1>
        <Canvas>
          <Thing />
        </Canvas>
      </Content>
    );
  }
}
export default Home;