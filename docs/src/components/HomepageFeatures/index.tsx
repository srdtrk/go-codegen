import React from "react";

import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

import EasyDeploySvg from "@site/static/img/easy_deploy.svg";
import UniversalSupportSvg from "@site/static/img/universal_support.svg";
import FocusSvg from "@site/static/img/focus.svg";

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: JSX.Element;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Generate Message Types',
    Svg: EasyDeploySvg,
    description: (
      <>
        go-codegen generates message types in golang for you CosmWasm application. It is as easy as
        running a command.
      </>
    ),
  },
  {
    title: 'Generate gRPC Clients',
    Svg: UniversalSupportSvg,
    description: (
      <>
        go-codegen generates gRPC clients in golang for you CosmWasm application. Currently, only the
        gRPC query client is supported. A transaction client is on the roadmap.
      </>
    ),
  },
  {
    title: 'A Streamlined IBC Testing Suite',
    Svg: FocusSvg,
    description: (
      <>
        go-codegen generates an entire testing suite for your CosmWasm application. This end-to-end
        testing suite is designed to help you test your application in a local and realistic environment.
        Powered by <a href="https://github.com/strangelove-ventures/interchaintest">interchaintest</a>.
      </>
    ),
  },
];

function Feature({ title, Svg, description }: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
