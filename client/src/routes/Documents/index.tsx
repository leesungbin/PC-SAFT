import React from 'react';
import { Content } from '../../components/Content';
import './index.css';
import './github.css';
import ReactHtmlParser from 'react-html-parser';
import { GITHUB_MARKDOWN_ENDPOINT } from '../../_lib/endpoint';
import { LinearProgress } from '@material-ui/core';

type State = {
  markdown?: string
}
class Document extends React.Component<{}, State> {
  state: State = {}
  componentDidMount = async () => {
    const res = await fetch('/Readme.md');
    const markdownRaw = await res.text();
    const githubRes = await fetch(GITHUB_MARKDOWN_ENDPOINT, {
      method: 'POST',
      body: JSON.stringify({ text: markdownRaw })
    })
    const markdown = await githubRes.text();
    this.setState({ markdown });
  }
  render() {
    const { markdown } = this.state;
    return (
      <>
        {!markdown && <div><LinearProgress /></div>}
        <Content>

          <article className="markdown-body" style={{ marginTop: 20 }}>
            {ReactHtmlParser(markdown ? markdown : '')}
          </article>
        </Content>
      </>
    );
  }
}
export default Document;