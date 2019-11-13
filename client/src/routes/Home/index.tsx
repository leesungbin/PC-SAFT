import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';
import { EQUIL_ENDPOINT } from '../../_lib/endpoint';
import Point from '../../threeFragments/Point';
import { TieLine } from '../../threeFragments/TieLine';

type HomeProps = {
  width: number,
  height: number,
};

type State = {
  data: {
    P: number,
    T: number,
    x: number[],
    y: number[],
  }[],
  waiting: boolean,
}
class Home extends React.Component<HomeProps, State> {
  state: State = {
    data: [],
    waiting: false,
  }

  callEquil = async () => {
    this.setState({ waiting: true });
    const res = await fetch(EQUIL_ENDPOINT, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ T: 300, id: [18, 35, 62] }),
    });
    const json = await res.json()
    this.setState({ data: json.data, waiting: false });
  }

  render() {
    const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];

    const { data, waiting } = this.state;
    const len = data.length;
    return (
      <div>
        <Content>
          <div style={{ height: 70 }}>
            <button onClick={() => this.callEquil()}>fetch test</button>
            <p>현재 물질 : 1-propylamine (N-PROPYL AMINE) / benzene / isobutane</p>
            {waiting && <p>계산 중이에요.</p>}
          </div>
        </Content>
        <Canvas
          style={{ height: this.props.height * 0.8, width: '100%' }}
          camera={{ position: [1 / 2, Math.sqrt(3) / 4, 50], fov: 2, near: 1, far: -1 }}
        >
          <mesh rotation={[0, 0, 0]}>
            <Triangle points={trianglePoints} />
            {len && data.map((e, i) => (
              <mesh key={i}>
                {Math.floor(Math.random() * 5) === 2 && (
                  <>
                    <Point abc={e.x} val={0} t={0} />
                    <Point abc={e.y} val={0} t={1} />
                    <TieLine x={e.x} y={e.y} val={0} color="green" />
                  </>
                )}
              </mesh>
            ))}
          </mesh>
        </Canvas>
      </div>
    );
  }
}

export default Home;