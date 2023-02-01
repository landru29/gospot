import { getNextKeyDef } from "@testing-library/user-event/dist/keyboard/getNextKeyDef";

export class Album {
    public name: string = '';
    public label: string = '';
    public images: Image[] = [];


    constructor(
        data?: any,
    ) {
        Object.assign(this, data);

        this.images = data.images.map((image: any) => new Image(image))
    }
}

export class Image {
    public url: string = '';

    constructor(
        data?: any,
    ) {
        Object.assign(this, data);
    } 
}