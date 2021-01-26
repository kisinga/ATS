import {Pipe, PipeTransform} from "@angular/core";
import {ObjectID} from "bson";

@Pipe({
  name: "dateFromObjectId",
})
export class DateFromObjectIdPipe implements PipeTransform {
  transform(objectId: string, ...args: unknown[]): unknown {
    return ObjectID.createFromHexString(objectId).getTimestamp();
  }
}
