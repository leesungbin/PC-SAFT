import React from 'react';

type lp = {
  lang: string,
  onClick: (lang: string) => void,
}
export function LangSelector({ lang, onClick }: lp) {
  return (
    <button style={styles.font} onClick={() => onClick(lang)}>{lang}</button>
  );
}

const styles: { [key: string]: React.CSSProperties } = {
  font: {
    marginTop: 10,
    fontWeight: "bold",
  }
};