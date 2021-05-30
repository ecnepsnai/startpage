import * as React from 'react';
import { Palette } from './Palette';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import {
    faShoppingCart,
    faSpinner,
} from '@fortawesome/free-solid-svg-icons';
import '../../css/Icon.scss';

export namespace Icon {
    interface IconProps {
        pulse?: boolean;
        spin?: boolean;
        color?: Palette;
    }

    interface EIconProps {
        icon: IconProp;
        options: IconProps;
    }

    export const EIcon: React.FC<EIconProps> = (props: EIconProps) => {
        let className = '';
        if (props.options.color) {
            className = 'color-' + props.options.color.toString();
        }
        return (<FontAwesomeIcon icon={props.icon} pulse={props.options.pulse} spin={props.options.spin} className={className} />);
    };

    interface LabelProps { icon: JSX.Element; spin?: boolean; label: string | number; }
    export const Label: React.FC<LabelProps> = (props: LabelProps) => {
        return (
            <span className="icon-label">
                { props.icon}
                <span className="label-text">{props.label}</span>
            </span>
        );
    };

    export const ShoppingCart: React.FC<IconProps> = (props: IconProps) => EIcon({ icon: faShoppingCart, options: props });
    export const Spinner: React.FC<IconProps> = (props: IconProps) => EIcon({ icon: faSpinner, options: props });
}
