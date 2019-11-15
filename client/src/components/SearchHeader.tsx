import React from 'react';
import SearchIcon from '@material-ui/icons/Search';
import { TextField, InputAdornment } from '@material-ui/core';

type SearchHeaderProps = {
  text: string,
  onChangeContent?: (text: string) => void,
  onSearched?: () => void,
  listComponents?: React.ReactElement,
}
export default function SearchHeader({ text, onChangeContent, listComponents }: SearchHeaderProps) {
  return (
    <div style={{ margin: '0, 9%, 0, 9%' }}>
      <TextField
        value={text}
        onChange={event => onChangeContent && onChangeContent(event.target.value)}
        placeholder="Search Components"
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <SearchIcon />
            </InputAdornment>
          )
        }}
        style={{ flex: 1, width: '80%', marginTop: 10, marginBottom: 10, paddingLeft: '10%', paddingRight: '10%' }}
      />
      <div style={{ position: 'relative', flexDirection: 'row', justifyContent: 'center' }}>
        {listComponents}
      </div>
    </div>
  )
}