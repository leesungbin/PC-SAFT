import React from 'react';
import { DATA_ENDPOINT } from '../../_lib/endpoint';
import DataTable from './Table';
import './index.css';
import { Select, MenuItem, InputLabel, FormControl } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

type FilterValue = 'Zero' | 'NonZero' | 'None';
type Component = {
  id: string,
  name: string,
  d: number,
  e: number,
}
type State = {
  columns: { title: string, field: string }[],
  data: Component[],
  dStatus: FilterValue,
  eStatus: FilterValue,
}

class Database extends React.Component<{}, State> {
  state: State = {
    columns: [
      { title: 'Name', field: 'name' },
      { title: 'Mw', field: 'Mw' },
      { title: 'Pc', field: 'Pc' },
      { title: 'Tc', field: 'Tc' },
      { title: 'Tb', field: 'Tb' },
      { title: 'ω', field: 'w' },
      { title: 'ε', field: 'epsilon' },
      { title: 'm', field: 'm' },
      { title: 'σ', field: 'sigma' },
      { title: 'k', field: 'k' },
      { title: 'e', field: 'e' },
      { title: 'd', field: 'd' },
      { title: 'x', field: 'x' },
    ],
    data: [],
    dStatus: 'None',
    eStatus: 'None',
  }

  componentDidMount = async () => {
    const fetchData = await fetch(DATA_ENDPOINT, { method: 'POST' });
    const json = await fetchData.json();
    const data = json.data.map((e: any) => { return e.data });

    this.setState({ data });
  }

  filterData(data: Component[]) {
    let results = data.slice();
    if (this.state.dStatus === 'NonZero') {
      results = results.filter((datum: Component) => datum.d !== 0)
    } else if (this.state.dStatus === 'Zero') {
      results = results.filter((datum: Component) => datum.d === 0)
    }
    if (this.state.eStatus === 'NonZero') {
      results = results.filter((datum: Component) => datum.e !== 0)
    } else if (this.state.eStatus === 'Zero') {
      results = results.filter((datum: Component) => datum.e === 0)
    }
    return results
  }

  onDStatusChange(e: any) {
    this.setState({ dStatus: e.target.value })
  }

  onTStatusChange(e: any) {
    this.setState({ eStatus: e.target.value })
  }

  render() {
    return (
      <div style={{ margin: 20, maxWidth: '100%' }}>
        <FormControl>
          <InputLabel id="d-status-select-label">d 필터</InputLabel>
          <Select 
            labelId="d-status-select-label"
            id="d-status-select"
            color="primary"
            value={this.state.dStatus} onChange={(e) => this.onDStatusChange(e)}>
            <MenuItem value={"None"}>필터 없음</MenuItem>
            <MenuItem value={"Zero"}>d가 0인 것만 보기</MenuItem>
            <MenuItem value={"NonZero"}>d가 0이 아닌 것만 보기</MenuItem>
          </Select>
        </FormControl>
        <FormControl>
          <InputLabel id="t-status-select-label">e 필터</InputLabel>
          <Select 
            labelId="t-status-select-label"
            id="t-status-select"
            color="primary"
            value={this.state.eStatus} onChange={(e) => this.onTStatusChange(e)}>
            <MenuItem value={"None"}>필터 없음</MenuItem>
            <MenuItem value={"Zero"}>e가 0인 것만 보기</MenuItem>
            <MenuItem value={"NonZero"}>e가 0이 아닌 것만 보기</MenuItem>
          </Select>
        </FormControl>
        <DataTable
          columns={this.state.columns}
          data={this.filterData(this.state.data)}
          title="DB Table"
        />
      </div>
    );
  }
}
export default Database;