import { Application } from '@hotwired/stimulus';
import LiveController from '@symfony/ux-live-component';

const app = Application.start();

// Register the live component controller
app.register('live', LiveController);

// Register any custom controllers in assets/controllers/
const controllers = import.meta.glob('./controllers/*.js', { eager: true });
for (const path in controllers) {
    const name = path.match(/\/([^/]+)_controller\.js$/)?.[1].replace(/_/g, '-');
    if (name) {
        app.register(name, controllers[path].default);
    }
}

export { app };
