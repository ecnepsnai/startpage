/**
 * Class for interacting with the API
 */
export class API {
    private static do(url: string, options: RequestInit, noCheckError?: boolean): Promise<unknown> {
        return new Promise((resolve, reject) => {
            fetch(url, options).then(response => {
                response.json().then(results => {
                    if (noCheckError) {
                        resolve(results.data);
                        return;
                    }

                    if (!response.ok) {
                        let message = 'Internal Server Error';
                        if (results.error && results.error.message) {
                            message = results.error.message;
                        }
                        console.error('API error caught', results, message);
                        reject(results);
                    } else {
                        resolve(results.data);
                    }
                }, e => {
                    if (!noCheckError) {
                        console.error('API error caught', e);
                    }
                    reject(e);
                });
            }, e => {
                if (!noCheckError) {
                    console.error('API error caught', e);
                }
                reject(e);
            });
        });
    }

    /**
     * Perform a HTTP GET request to the specified URL
     * @param url the URL to request
     * @returns The JSON object of the results
     */
    public static GET(url: string): Promise<unknown> {
        return this.do(url, { method: 'GET' });
    }

    /**
     * Perform a HTTP POST request to the specified URL with the given body
     * @param url the URL to request
     * @param data Body data to be encoded as JSON
     * @returns The JSON object of the results
     */
    public static async POST(url: string, data: unknown): Promise<unknown> {
        return this.do(url, { method: 'POST', body: JSON.stringify(data) });
    }

    /**
     * Perform a HTTP PUT request to the specified URL with the given body
     * @param url the URL to request
     * @param data Body data to be encoded as JSON
     * @returns The JSON object of the results
     */
    public static async PUT(url: string, data: unknown): Promise<unknown> {
        return this.do(url, { method: 'PUT', body: JSON.stringify(data) });
    }

    /**
     * Perform a HTTP PATCH request to the specified URL with the given body
     * @param url the URL to request
     * @param data Body data to be encoded as JSON
     * @returns The JSON object of the results
     */
    public static async PATCH(url: string, data: unknown): Promise<unknown> {
        return this.do(url, { method: 'PATCH', body: JSON.stringify(data) });
    }

    /**
     * Perform a HTTP DELETE request to the specified URL
     * @param url the URL to request
     * @returns The JSON object of the results
     */
    public static async DELETE(url: string): Promise<unknown> {
        return this.do(url, { method: 'DELETE' });
    }
}
