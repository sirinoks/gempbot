import React, { StrictMode } from 'react';
import ReactDOM from 'react-dom';
import { createGlobalStyle } from 'styled-components';
import { App } from './App';

const GlobalStyle = createGlobalStyle`
    body {
        --bg: #0e0e10;
        --bg-bright: #18181b;
        --bg-brighter: #3d4146;
        --bg-dark: #121416;
        --theme: #009148;
        --theme-bright: #00a552;
        --theme2: #2980b9;
        --theme2-bright: #3498db;
        --theme2-dark: #24618a;
        --text: #F5F5F5;
        --text-dark: #616161;
        --twitch: #6441a5;
        --twitch-dark: #4c317e;
        --danger: #e74c3c;
        --danger-dark: #c0392b;

        background: var(--bg);
        margin: 0;
        padding: 0;
        color: var(--text);
        margin: 0;
        font-family: Helvetica, Arial, sans-serif;
        height: 100%;
        width: 100%;

        * {
            box-sizing: border-box;
        }
    }
`

ReactDOM.render(
    <StrictMode>
        <React.Fragment>
            <GlobalStyle />
            <App />
        </React.Fragment>
    </StrictMode>,
    document.getElementById('root')
);