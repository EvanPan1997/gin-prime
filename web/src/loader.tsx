import React, {Suspense} from "react";

export function loadComponent(componentPath: string, componentName: string) {
    const LazyComponent = React.lazy(() => import(`./${componentPath}/${componentName}.tsx`));
    if (!LazyComponent) {
        return null
    }
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <LazyComponent/>
        </Suspense>
    )
}