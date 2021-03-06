import * as React from 'react';

interface ErrorBoundaryState {
    errorDetails?: string;
}
export class ErrorBoundary extends React.Component<unknown, ErrorBoundaryState> {
    constructor(props: unknown) {
        super(props);
        this.state = {};
    }

    static getDerivedStateFromError(error: Error): ErrorBoundaryState {
        const errData = JSON.stringify(error, Object.getOwnPropertyNames(error), 4);
        console.log(errData);
        return { errorDetails: errData };
    }

    componentDidCatch(error: Error, errorInfo: React.ErrorInfo): void {
        console.error('An error occurred: ' + error, errorInfo);
    }

    render(): JSX.Element {
        if (this.state.errorDetails) {
            // Remove any modal backdrops in case the exception occurred in a modal
            return (
                <div className="container">
                    <div className="card mt-3">
                        <div className="card-header">
                            An Error Occurred
                        </div>
                        <div className="card-body">
                            <p>An unrecoverable error occurred while attempting to render this page.
                                Please report this as an issue on <a href="https://github.com/ecnepsnai/furdl/issues/new/choose" target="_blank" rel="noreferrer">Github</a> and include the following information:</p>
                            <pre>{this.state.errorDetails}</pre>
                        </div>
                    </div>
                </div>
            );
        }

        return (<React.Fragment>{this.props.children}</React.Fragment>);
    }
}
