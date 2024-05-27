import React from 'react';
import CodeBlock from '@theme-original/CodeBlock';
import type CodeBlockType from '@theme/CodeBlock';
import type { WrapperProps } from '@docusaurus/types';

type Props = WrapperProps<typeof CodeBlockType> & {
  source?: string;
}

export default function CodeBlockWrapper(props: Props): JSX.Element {
  const { source, ...codeBlockProps } = props;
  return (
    <>
      <CodeBlock {...codeBlockProps} />
      {source &&
        <div className='text-right mb-4'>
          <a href={source} target="_blank" rel="noopener noreferrer">View Source</a>
        </div>
      }
    </>
  );
}
