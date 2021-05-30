import * as React from 'react';
import '../../css/Card.scss';

export interface CardProps {
    header: JSX.Element;
    children?: React.ReactNode;
}
export const Card: React.FC<CardProps> = (props: CardProps) => {
    return (
        <div className="card">
            <div className="card-head">{props.header}</div>
            <div className="card-body">{props.children}</div>
        </div>
    );
};
