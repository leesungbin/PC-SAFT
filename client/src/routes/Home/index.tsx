import React from 'react';
import { Content } from '../../components/Content';
import { EQUIL_ENDPOINT, DATA_ENDPOINT, FLASHES_ENDPOINT } from '../../_lib/endpoint';


import ContinuosSlider from '../../components/ContinuosSlider';
import { Typography, Chip, FormGroup, FormControlLabel, TextField } from '@material-ui/core';
import Switch from '@material-ui/core/Switch';
import CalculatingIndicator from '../../components/CalculatingIndicator';
import { ErrorSnack } from '../../components/Snack';

import './index.css';
import ComponentSelector from '../../components/ComponentSelector';
import FormControlCondition from '../../components/FormControlCondition';
import { Stage, Layer, Line, Text } from 'react-konva';
import Tie from '../../canvas/Tie';
import { xyTransform } from '../../_lib/coordChange';
import { CLs } from '../../canvas/CL';
import Point from '../../canvas/Point';

type FetchResult = {
  result: {
    data: { P: number, T: number, x: number[], y: number[] }[],
    header: { min: number, max: number },
    mode: string,
    names: string[],
  }
}
type HomeProps = {
  width: number,
  height: number,
}
export type Component = {
  id: string,
  name: string,
  selected?: boolean,
}
type State = {
  data: {
    P: number,
    T: number,
    x: number[],
    y: number[],
  }[],
  waiting: boolean,
  T?: number,
  P?: number,
  fetchT?: number,
  fetchP?: number,
  openSelector: boolean,
  mode?: string,
  min?: number,
  max?: number,
  names?: string[],
  selectedComponents: Component[],
  error: string,
  components: Component[],
  constT: boolean,
  constP: boolean,
  showcl: boolean,
  showtie: boolean,
  showexp: boolean,
  expData?: number[][],
  expText?: string,
}

class Home extends React.Component<HomeProps, State> {
  state: State = {
    data: [],
    waiting: false,
    // T: 300,
    // P: 1,
    openSelector: false,
    selectedComponents: [],
    error: '',
    components: [],
    constT: false,
    constP: false,
    showcl: true,
    showtie: false,
    showexp: false,
  }

  componentDidMount = async () => {
    const fetchData = await fetch(DATA_ENDPOINT, { method: 'POST' });
    const json = await fetchData.json();
    const components = json.data.map((e: any) => { return { ...e.data, id: e.id } });

    this.setState({ components });
  }
  componentDidUpdate = (_: HomeProps, prevState: State) => {
    if (prevState.components !== this.state.components) {
      const selectedComponents = this.state.components.filter(component => component.selected === true);
      this.setState({ selectedComponents });
    }
    // console.log(this.state.T, this.state.P);
  }
  startCalculate = () => {
    if (this.state.error) return;
    if (this.state.selectedComponents.length !== 3) {
      this.setState({ error: '3개의 물질을 선택하세요.' });
      return;
    }
    if (!this.state.constT && !this.state.constP) {
      this.setState({ error: '온도나 압력을 고정해주세요.' });
      return;
    }
    if (this.state.constT && this.state.constP) {
      if (!this.state.fetchT && !this.state.fetchP) {
        this.setState({ error: 'Constant 설정한 property에 값을 입력해야합니다.' });
        return
      }
      this.callFlashes();
      return;
    }
    if ((this.state.constT && this.state.fetchT) || (this.state.constP && this.state.fetchP)) {
      console.log('call equil');
      this.callEquil();
    }
  }
  callEquil = async () => {
    const { selectedComponents } = this.state;
    console.log(selectedComponents);
    const id = selectedComponents.map(comp => { return comp.id });

    this.setState({ waiting: true });
    const res = await fetch(EQUIL_ENDPOINT, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: this.state.constP ? JSON.stringify({ P: this.state.fetchP, id }) : JSON.stringify({ T: this.state.fetchT, id }),
    });
    const json: FetchResult = await res.json()
    const { data, header, mode, names } = json.result;
    if (!data || data.length === 0) {
      this.setState({ error: '계산이 잘 되지 않는 조합입니다...', waiting: false });
      return;
    }
    const { min, max } = header;
    this.setState({ data, min, max, mode, names, waiting: false });
    if (mode === 'BUBLP') {
      this.setState({ P: min });
    } else {
      this.setState({ T: min });
    }
  }

  // fetch method 작성할 것.
  callFlashes = async () => {
    const { selectedComponents } = this.state;
    console.log(selectedComponents);
    const id = selectedComponents.map(comp => { return comp.id });

    this.setState({ waiting: true });
    const res = await fetch(FLASHES_ENDPOINT, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ P: this.state.fetchP, T: this.state.fetchT, id }),
    });
    const json: FetchResult = await res.json();
    console.log(json);
    const { data, names } = json.result;
    const mode = 'FLASH';

    if (!data || data.length === 0) {
      this.setState({ error: '계산이 잘 되지 않는 조합입니다...', waiting: false, mode });
      return;
    }
    this.setState({ data, mode, names, waiting: false });
  }

  cancleComponent = (index: number) => {
    let components = [...this.state.components]
    const cancleComponentIdx = this.state.components.findIndex(component => component.name === this.state.selectedComponents[index].name);
    if (cancleComponentIdx > -1) {
      components[cancleComponentIdx].selected = false;
      this.setState({ components });
    }
  }

  saveCanvas = () => {
    const canvasSave = this.refs.canvas as HTMLCanvasElement;
    const d = canvasSave && canvasSave.toDataURL();
    const w = window.open();
    w && w.document.write("<img src='" + d + "' alt='img from canvas'/>");
  }

  setExpData = (expText: string) => {
    // const { expText } = this.state;
    this.setState({ expText })
    let expData: number[][] = [];
    if (expText) {
      const lines: string[] = expText.split('\n');
      lines.forEach(e => {
        let line: number[] = [];
        const ok = e.split(' ').length === 3;
        if (ok) {
          e.split(' ').forEach((k => {
            line.push(parseFloat(k));
          }));
          expData.push(line);
        }
      });
    }
    this.setState({expData});
  }
  render() {
    // const trianglePoints = [new Vector3(0, 0, 0), new Vector3(1, 0, 0), new Vector3(1 / 2, Math.sqrt(3) / 2, 0)];

    const { data, waiting, T, P, mode, names } = this.state;
    // const len = data.length;
    const { width } = this.props;
    const isMobile = (width) < 850 ? true : false;
    const ternaryWidth = isMobile ? width - 20 : Math.min((width - 380) * 0.7 - 20, 650);
    const points = [ternaryWidth * 0.015026, ternaryWidth * 0.92, ternaryWidth * 0.5, ternaryWidth * 0.08, ternaryWidth * 0.984974, ternaryWidth * 0.92];

    const translated = data.map(e => {
      // const xy_x = coordChange(e.x[0], e.x[1], e.x[2]);
      // const xy_y = coordChange(e.y[0], e.y[1], e.y[2]);
      const liq = xyTransform(points, e.x);
      const vap = xyTransform(points, e.y);
      return { liq, vap, T: e.T, P: e.P, x: e.x, y: e.y };
    });

    const expData = this.state.expData && this.state.expData.map(e => {
      return xyTransform(points, e)
    });

    return (
      <div>
        <ul style={isMobile
          ? { marginLeft: '10%', padding: 0, marginTop: 10, width: '80%', overflowX: 'scroll', height: 40, whiteSpace: 'nowrap' }
          : { marginTop: 10, marginLeft: "7%", height: 50, whiteSpace: 'nowrap', padding: 0, overflowX: 'scroll', width: '86%' }}>
          {this.state.selectedComponents && this.state.selectedComponents.map((comp, i) => (
            <li key={i} style={{ display: 'inline', listStyle: 'none', margin: 0, padding: 0 }}>
              <Chip style={{ marginRight: 10, marginBottom: 10 }} label={comp.name} variant="outlined" onDelete={() => this.cancleComponent(i)} />
            </li>
          ))}
        </ul>
        <div style={{ display: 'flex', justifyContent: 'space-around', zIndex: 0, flexWrap: "wrap" }}>
          <div style={isMobile ? { width: '80%', margin: '10%', marginTop: 10 } : { width: 300, padding: 30, paddingTop: 10 }}>
            <FormGroup>
              <FormControlCondition valueDef="P" placeholder="Const Pressure (atm)" onChangeValue={(fetchP) => this.setState({ fetchP })} onError={(error) => this.setState({ error })} onCheckConst={(constP) => this.setState({ constP })} />
              <FormControlCondition valueDef="T" placeholder="Const Temperature (K)" onChangeValue={(fetchT) => this.setState({ fetchT })} onError={(error) => this.setState({ error })} onCheckConst={(constT) => this.setState({ constT })} />
            </FormGroup>

            <ComponentSelector
              components={this.state.components}
              onClickContent={(i) => {
                let newComponents = [...this.state.components];
                // switching value
                newComponents[i].selected = newComponents[i].selected ? false : true;
                this.setState({ components: newComponents });
              }}
              calcButton={
                <button onClick={() => { this.startCalculate(); }} style={styles.calculateButton}>Calculate
              </button>}
            />
            <button onClick={() => this.saveCanvas()} style={styles.saveButton}>Save Image</button>

            {/* exp data input */}
          </div>
          <div style={{ position: 'relative' }}>
            <div style={{ position: 'absolute', top: 0, left: 10, zIndex: 20 }}>
              <FormGroup>
                <FormControlLabel
                  control={<Switch checked={this.state.showcl} onChange={() => this.setState({ showcl: !this.state.showcl })} value="Show connected line" />}
                  label="Binodal curve"
                />
                <FormControlLabel
                  control={<Switch checked={this.state.showtie} onChange={() => this.setState({ showtie: !this.state.showtie })} value="Show tieline" />}
                  label="Tieline"
                />
                <FormControlLabel
                  control={<Switch checked={this.state.showexp} onChange={() => this.setState({ showexp: !this.state.showexp })} value="Show exp data" />}
                  label="Exp data"
                />
              </FormGroup>
            </div>
            <Stage ref='canvas' width={ternaryWidth} height={ternaryWidth} style={{ zIndex: 10, marginRight: 10, marginLeft: 10, marginTop: -30 }}>
              <Layer>
                {/* label */}
                <Text text={names && names[0]} x={0} y={ternaryWidth * 0.08 - 21} width={ternaryWidth} align="center" fontSize={20} />
                <Text text={names && names[1]} x={0} y={ternaryWidth * 0.92 + 1} fontSize={20} />
                <Text text={names && names[2]} x={0} y={ternaryWidth * 0.92 + 1} width={ternaryWidth} align="right" fontSize={20} />

                {/* background triangle */}
                <Line points={points} closed={true} stroke="black" strokeWidth={3} />

                {/* fill with calculated data */}
                {/* bublP */}
                {mode === "BUBLP" && data && P && translated.map((e, i) => {
                  if (P < e.P * 1.001 && P > e.P * 0.999) {
                    return <Tie key={i} liq={e.liq} vap={e.vap} info={{ L: e.x, V: e.y }} showtie={this.state.showtie} />
                  }
                  return null;
                })}

                {/* bublT */}
                {mode === "BUBLT" && data && T && translated.map((e, i) => {
                  if (T < e.T * 1.001 && T > e.T * 0.999) {
                    return <Tie key={i} liq={e.liq} vap={e.vap} info={{ L: e.x, V: e.y }} showtie={this.state.showtie} />
                  }
                  return null;
                })}

                {/* for flash */}
                {mode === "FLASH" && data && translated.map((e, i) => {
                  return <Tie key={i} liq={e.liq} vap={e.vap} info={{ L: e.x, V: e.y }} showtie={this.state.showtie} />
                })}
                {mode === "FLASH" && data && this.state.showcl && <CLs datas={translated} />}

                {/* exp data plot */}
                {expData && this.state.showexp &&
                  expData.map((e, i) =>
                    <Point key={i} hover={false} setHover={() => { }} fill="purple" x={e.x} y={e.y} info={[]} />
                  )
                }
              </Layer>
            </Stage>
            <Content>
              <div style={{ height: 100, justifyContent: 'center' }}>
                <div style={{ display: 'flex', width: '100%', flexDirection: 'row', justifyContent: 'space-around' }}>
                  {P && mode && mode === "BUBLP" ?
                    <div style={{ flex: 1, padding: '0,10%,10%,0' }}>
                      <Typography gutterBottom>P : {P.toFixed(3)} atm</Typography>
                      <ContinuosSlider val={P} onChange={(P) => this.setState({ P })} min={this.state.min} max={this.state.max} />
                    </div>
                    : T && mode === "BUBLT" ?
                      <div style={{ flex: 1, padding: '0,10%,10%,0' }}>
                        <Typography gutterBottom>T : {T.toFixed(3)} K</Typography>
                        <ContinuosSlider val={T} onChange={(T) => this.setState({ T })} min={this.state.min} max={this.state.max} />
                      </div>
                      : <></>
                  }
                </div>
              </div>
              {this.state.showexp && <TextField style={{width: '100%', marginBottom: 20 }} multiline placeholder="exp data (ex:  0.2 0.3 0.5)" value={this.state.expText} onChange={(e) => this.setExpData(e.target.value)} />}
            </Content>
          </div>
        </div>
        <CalculatingIndicator open={waiting} />
        <ErrorSnack error={this.state.error} onClose={() => this.setState({ error: '' })} />
      </div >
    );
  }
}
const styles: { [key: string]: React.CSSProperties } = {
  calculateButton: {
    width: '100%', height: 40, backgroundColor: "#FECF58", borderRadius: 10, fontSize: 17, border: 'none'
  },
  saveButton: {
    marginTop: 10, width: '100%', height: 40, backgroundColor: "#ABCDEF", borderRadius: 10, fontSize: 17, border: 'none',
  }
}
// const mobileStyles: { [key: string]: React.CSSProperties } = {

// }

export default Home;