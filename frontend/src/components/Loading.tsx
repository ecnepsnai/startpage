import * as React from 'react';
import { Icon } from './Icon';
import '../../css/Loading.scss';

export const Loading: React.FC = () => {
    return (
        <div className="loading">
            <Icon.Spinner pulse />
            <span className="loading-text">Loading...</span>
        </div>
    );
};
