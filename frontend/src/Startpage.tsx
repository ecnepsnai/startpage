import * as React from 'react';
import '../css/Startpage.scss';
import { MEDailyDealCard } from './modules/medailydeal/MEDailyDealCard';
import { WeatherCard } from './modules/weather/WeatherCard';

export const Startpage: React.FC = () => {
    const spStyle: React.CSSProperties = {
        backgroundImage: 'url("/api/modules/potd/picture")',
    };
    return (
        <div className="startpage" style={spStyle}>
            <div className="modules">
                <MEDailyDealCard />
                <WeatherCard />
            </div>
        </div>
    );
};
