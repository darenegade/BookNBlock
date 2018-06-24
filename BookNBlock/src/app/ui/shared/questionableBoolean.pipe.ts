import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ name: 'questionableBoolean' })
export class QuestionableBooleanPipe implements PipeTransform {
    transform(value: boolean): string {
        return value === true ? 'Yes' : 'No';
    }
}
