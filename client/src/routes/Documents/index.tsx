import React from 'react';
import { Content } from '../../components/Content';
import './index.css';
import './github.css';
import ReactHtmlParser from 'react-html-parser';
import { GITHUB_MARKDOWN_ENDPOINT } from '../../_lib/endpoint';
import { LinearProgress } from '@material-ui/core';
import { LangSelector } from './LangSelector';

type State = {
  markdown?: string,
  loading: boolean,
  lang: string,
}
class Document extends React.Component<{}, State> {
  state: State = { lang: 'ko', loading: true };
  changeLang = (lang: string) => {
    this.setState({ lang, loading: true });
    this.loadMD(lang);
  }
  loadMD = async (lang: string) => {
    let path: string;
    if (lang === 'ko') {
      path = '/Readme.md';
    } else {
      path = '/Readme_eng.md';
    }
    const res = await fetch(path);
    const markdownRaw = await res.text();
    const githubRes = await fetch(GITHUB_MARKDOWN_ENDPOINT, {
      method: 'POST',
      body: JSON.stringify({ text: markdownRaw })
    })
    const markdown = await githubRes.text();
    this.setState({ markdown, loading: false });
  }
  componentDidMount = () => {
    this.loadMD('ko');
  }
  render() {
    const { markdown, loading } = this.state;
    return (
      <>
        {loading && <div className="sticky"><LinearProgress /></div>}
        <Content>
          <div style={{ display: 'flex' }}>
            <LangSelector lang="ko" onClick={this.changeLang} />
            <div style={{ width: 10 }}></div>
            <LangSelector lang="en" onClick={this.changeLang} />
          </div>
          <article className="markdown-body" style={{ marginTop: 20, marginBottom: 20 }}>
            {ReactHtmlParser(markdown ? markdown : '')}
          </article>
        </Content>
      </>
    );
  }
}
export default Document;