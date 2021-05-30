import * as React from 'react';
import { Card } from '../../components/Card';
import { Icon } from '../../components/Icon';
import { Loading } from '../../components/Loading';
import { API } from '../../services/API';
import { MEDailyDeal } from '../../types/MEDailyDeal';

export const MEDailyDealCard: React.FC = () => {
    const [IsLoading, setIsLoading] = React.useState(true);
    const [Deal, setDeal] = React.useState<MEDailyDeal>();

    const getDeal = () => {
        API.GET('/api/modules/medailydeal/info').then(data => {
            setDeal(data as MEDailyDeal);
            setIsLoading(false);
        });
    };

    React.useEffect(() => {
        getDeal();
    }, []);

    return (
        <Card header={<Icon.Label icon={<Icon.ShoppingCart />} label="MemoryExpress Deal of the Day" />}>
            {IsLoading ? (<Loading />) : (<DailyDealBody deal={Deal} />)}
        </Card>
    );
};

interface DailyDealBodyProps {
    deal: MEDailyDeal;
}
const DailyDealBody: React.FC<DailyDealBodyProps> = (props: DailyDealBodyProps) => {
    return (
        <div className="me-daily-deal">{props.deal.Title}</div>
    );
};
