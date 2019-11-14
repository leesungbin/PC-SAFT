import React from 'react';
import { Content } from '../../components/Content';
import { Canvas } from 'react-three-fiber';
import { Triangle } from '../../threeFragments/Triangle';
import { Vector3 } from 'three';
import { EQUIL_ENDPOINT } from '../../_lib/endpoint';
import Point from '../../threeFragments/Point';
import { TieLine } from '../../threeFragments/TieLine';

import ContinuosSlider from '../../components/ContinuosSlider';
import { Typography } from '@material-ui/core';

type FetchResult = {
  result: {
    data: { P: number, T: number, x: number[], y: number[] }[],
    header: { min: number, max: number },
    mode: string
  }
}
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
  T: number,
  P: number,
  mode?: string,
  min?: number,
  max?: number,
}
class Home extends React.Component<HomeProps, State> {
  state: State = {
    data: [],
    waiting: false,
    T: 300,
    P: 1,
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
    const json: FetchResult = await res.json()
    const { data, header, mode } = json.result;
    const { min, max } = header;
    this.setState({ data, min, max, mode, waiting: false });
  }

  render() {
    const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];

    const { data, waiting, T, P, mode } = this.state;
    const len = data.length;

    const textPosition = new Vector3(1 / 2, Math.sqrt(3) / 4 + 0.2, 0);

    return (
      <div>
        <Content>
          <div style={{ height: 100 }}>
            <button onClick={() => this.callEquil()}>fetch test</button>
            {/* <p>현재 물질 : 1-propylamine (N-PROPYL AMINE) / benzene / isobutane</p> */}
            {waiting && <p>계산 중이에요.</p>}
          </div>
          <div style={{ display: 'flex', flexDirection: 'row', justifyContent: 'space-around' }}>
            {mode && mode === "BUBLP" ?
              <div style={{ flex: 1, padding: 10 }}>
                <Typography gutterBottom>P : {P.toFixed(3)} atm</Typography>
                <ContinuosSlider val={P} onChange={(P) => this.setState({ P })} min={this.state.min} max={this.state.max} />
              </div>
              : mode === "BUBLT" ?
                <div style={{ flex: 1, padding: 10 }}>
                  <Typography gutterBottom>T : {T.toFixed(3)} K</Typography>
                  <ContinuosSlider val={T} onChange={(T) => this.setState({ T })} min={this.state.min} max={this.state.max} />
                </div>
                : <></>
            }

          </div>
        </Content>
        <div style={{ display: 'flex', justifyContent: 'center', }}>
          <Canvas
            style={{ height: 500, width: 500 }}
            camera={{ position: [1 / 2, Math.sqrt(3) / 4, 50], fov: 2, near: 0.5, far: -0.2 }}
          >
            <mesh rotation={[0, 0, 0]}>
              <Triangle points={trianglePoints} />
              {len && data.map((e, i) => {
                if (mode === "BUBLP" && e.P < P * 1.03 && e.P > P * 0.97) {
                  return (
                    <mesh key={i}>
                      <Point abc={e.x} val={0} t={0} />
                      <TieLine x={e.x} y={e.y} val={0} color="green" />
                      <Point abc={e.y} val={0} t={1} />
                    </mesh>
                  )
                }
                else if (mode === "BUBLT" && e.T < T * 1.03 && e.T > T * 0.97) {
                  return (
                    <mesh key={i}>
                      <Point abc={e.x} val={0} t={0} />
                      <TieLine x={e.x} y={e.y} val={0} color="green" />
                      <Point abc={e.y} val={0} t={1} />
                    </mesh>
                  )
                }
              })}
            </mesh>
          </Canvas>
        </div>
      </div>
    );
  }
}

export default Home;