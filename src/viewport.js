//@flow
import { View, MouseEvent, Point, Layer, Project } from 'paper';


/**
 * Encapsulate the viewport's logic.
 * Expose an interface to define how the canvas viewed by this viewport instance
 * should be shown to the user.
 */
export default class Viewport {

  view: View;

  _mainLayer: Layer;
  _rasterLayer: Layer;
  _shouldRaster: bool;

  constructor(view: View, project: Project) {
    this.view = view;
    this.view.autoUpdate = false;

    this._mainLayer = project.activeLayer;
    this._rasterLayer = project.addLayer(new Layer());
    this._shouldRaster = false;

    this.view.onMouseDrag = (e) => this._translate(e);
    this.view.onMouseUp = (e) => {
      this._shouldRaster = false;
      this.redraw();
    }
  }

  /**
   * Set the viewport's zoom factor.
   */
  setZoom(value: number) {
    this.view.setZoom(value);
  }

  /**
   * Set the viewport's center at provided position.
   */
  setCenter(x: number, y: number) {
    this.view.center = new Point(x, y);
  }

  /**
   * redraw the viewport's display.
   */
  redraw() {
    requestAnimationFrame(() => {
      if (this._shouldRaster) {
        if(this._rasterLayer.isEmpty())
          this._rasterLayer.addChild(this._mainLayer.rasterize());

        // hide the main layer and activate the raster layer.
        this._mainLayer.visible = false;
        this._rasterLayer.visible = true;
      } else {
        // hide the raster layer and activate the main layer.
        this._rasterLayer.visible = false;
        this._mainLayer.visible = true;
      }

      this.view.update();
    });
  }

  /**
   * Implement how the viewport should translate.
   */
  _translate(e: MouseEvent) {
    const delta = e.delta;

    // reduce velocity of the drag
    // Since it's called multiple time per seconds and we need to divide the lenght of the
    // delta by 2, use shift operator for sick performances
    delta.length >>= 1;

    this.view.center = this.view.center.add(new Point(-delta.x, -delta.y));

    // Raster the main layer when translating the viewport
    this._shouldRaster = true;

    // redraw the viewport
    this.redraw();
  }
}
