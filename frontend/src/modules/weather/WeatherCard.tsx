import * as React from 'react';
import { Card } from '../../components/Card';
import { Icon } from '../../components/Icon';
import { Loading } from '../../components/Loading';
import { API } from '../../services/API';
import { Weather } from '../../types/Weather';

export const WeatherCard: React.FC = () => {
    const [IsLoading, setIsLoading] = React.useState(true);
    const [Weather, setWeather] = React.useState<Weather>();

    const getWeather = () => {
        API.GET('/api/modules/weather/info').then(data => {
            setWeather(data as Weather);
            setIsLoading(false);
        });
    };

    React.useEffect(() => {
        getWeather();
    }, []);

    return (
        <Card header={<Icon.Label icon={<Icon.ShoppingCart />} label="Weather" />}>
            {IsLoading ? (<Loading />) : (<WeatherBody weather={Weather} />)}
        </Card>
    );
};

interface WeatherBodyProps {
    weather: Weather;
}
const WeatherBody: React.FC<WeatherBodyProps> = (props: WeatherBodyProps) => {
    return (
        <div className="me-daily-weather">{props.weather.Current.Summary}</div>
    );
};
