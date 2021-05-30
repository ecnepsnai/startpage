import * as React from 'react';
import { ErrorBoundary } from './components/ErrorBoundary';
import { Startpage } from './Startpage';
import '../css/App.scss';

export const App: React.FC = () => {
    return (<ErrorBoundary><Startpage /></ErrorBoundary>);
};
