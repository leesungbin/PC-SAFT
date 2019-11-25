import React from 'react';
import { Component } from '../routes/Home';
import SearchHeader from './SearchHeader';
import { SEARCH_ENDPOINT } from '../_lib/endpoint';

type ComponentSelectorProps = {
  components: Component[],
  onClickContent: (idx: number) => void,
  calcButton: React.ReactElement
};
type State = {
  searchingName: string,
  searchingLists: any[],
}
export default class ComponentSelector extends React.Component<ComponentSelectorProps, State> {
  state: State = {
    searchingName: '',
    searchingLists: [],
  }
  callSearch = async (searchingName: string) => {
    const res = await fetch(SEARCH_ENDPOINT, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ name: searchingName.toLowerCase() })
    });
    const json = await res.json();
    this.setState({ searchingLists: json.data });
    // console.log(this.state.searchingLists);
  }

  render() {  
    const components = this.props.components.filter((component) => component.name.toLowerCase().indexOf(this.state.searchingName.toLowerCase()) > -1);
    return (
      <div>
        <SearchHeader
          text={this.state.searchingName}
          onChangeContent={(searchingName)=>{
            this.setState({searchingName});
            searchingName.length >=2 && this.callSearch(searchingName);
          }}
        />
        <ul style={{ height: 300, overflow: 'auto', padding: 0 }} className={'component-list'}>
          {components.map((component,i ) => (
            (component.name.indexOf(this.state.searchingName) > -1) ? <li key={i} style={{ backgroundColor: component.selected ? 'skyblue' : 'white' }}
              className={'component-list__item'}
              onClick={() => this.props.onClickContent(i)}>
              <span>{component.name}</span>
            </li> : <li key={i} style={{display: 'none'}}></li>
          ))}
        </ul>
        {this.props.calcButton}
      </div>
    )
  }
}