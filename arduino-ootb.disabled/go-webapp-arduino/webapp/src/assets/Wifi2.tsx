import * as React from "react";
import { SVGProps } from "react";

const SvgWifi = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={14}
    height={14}
    fill="none"
    viewBox="0 0 14 14"
    {...props}
  >
    <path fill="green" d="M6.99 12.243a1 1 0 1 0 0-2 1 1 0 0 0 0 2z" />
    <path
      fill="green"
      d="M9.11 9.622a.5.5 0 0 1-.35-.15 2.5 2.5 0 0 0-3.54 0 .5.5 0 0 1-.85-.353.5.5 0 0 1 .145-.352 3.5 3.5 0 0 1 4.95 0 .5.5 0 0 1-.355.855z"
    />
    <path
      fill="green"
      d="M11.235 7.497a.5.5 0 0 1-.355-.145 5.5 5.5 0 0 0-7.78 0 .5.5 0 0 1-.705-.705 6.5 6.5 0 0 1 9.19 0 .5.5 0 0 1 0 .705.5.5 0 0 1-.35.145z"
    />
    <path
      fill="green"
      d="M13.355 5.377a.5.5 0 0 1-.365-.135 8.5 8.5 0 0 0-12 0 .5.5 0 0 1-.85-.352.5.5 0 0 1 .145-.353 9.5 9.5 0 0 1 13.43 0 .5.5 0 0 1 0 .705.5.5 0 0 1-.36.135z"
      opacity={0.2}
    />
  </svg>
);

export default SvgWifi;

