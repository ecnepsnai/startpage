export interface Sample {
    Summary: string;
    Description: string;
    IconCode: string;
    Temp: number;
    TempHigh: number;
    TempLow: number;
}

export interface Forecast {
    Samples: Sample[];
}

export interface Weather {
    Current: Sample;
    Forecast: Forecast;
    Expires: string;
}
